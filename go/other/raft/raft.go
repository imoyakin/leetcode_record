package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"time"
)

//Node node storage
type node struct {
	address string
}

func newNode(address string) *node {
	return &node{
		address: address,
	}
}

// rafat status
const (
	follower = iota
	candidate
	leader
)

//Raft node
type Raft struct {
	id int

	state int

	currentTerm int

	votedFor int

	voteCount int

	heartbeatC chan bool

	toLeaderC chan bool

	nodes map[int]*node
}

func (r *Raft) start() {
	r.state = follower
	r.currentTerm = 0
	r.votedFor = -1
	r.heartbeatC = make(chan bool)
	r.toLeaderC = make(chan bool)

	go func() {
		rand.Seed(time.Now().UnixNano())

		for {
			switch r.state {
			case follower:
				select {
				case <-r.heartbeatC: //heartbeat
					log.Printf("follower -%d recived heartbeat\n", r.id)
				//avoid all node overtime at sametime
				case <-time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond): //heartbeat timeout
					log.Printf("follower-%d timeout\n", r.id)
					r.state = candidate
				}
			case candidate:
				fmt.Printf("Node: %d, now is candidate\n", r.id)

				r.currentTerm++
				r.votedFor = r.id
				r.voteCount = 1

				// broadcast i'm a candidate
				go r.broadcastRequestVote()

				select {
				case <-time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond):
					r.state = follower
				case <-r.toLeaderC:
					fmt.Printf("Node: %d, i'm leader\n", r.id)
					r.state = leader
				}
			case leader:
				r.broadcastHeartbeat()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

//VoteArgs request
type VoteArgs struct {
	Term        int
	CandidateID int
}

//VoteReply response
type VoteReply struct {
	Term        int
	VoteGranted bool
}

func (r *Raft) broadcastRequestVote() {
	var args = VoteArgs{
		Term:        r.currentTerm,
		CandidateID: r.id,
	}

	for i := range r.nodes {
		go func(i int) {
			var reply VoteReply
			r.sendRequestVote(i, args, &reply)
		}(i)
	}
}

func (r *Raft) sendRequestVote(serverID int, args VoteArgs, reply *VoteReply) {
	client, err := rpc.DialHTTP("tcp", r.nodes[serverID].address)
	if err != nil {
		log.Fatal(err)
	}

	client.Call("Raft.RequestVote", args, reply)

	if reply.Term > r.currentTerm {
		r.currentTerm = reply.Term
		r.state = follower
		r.votedFor = -1
		return
	}

	if reply.VoteGranted {
		r.voteCount++
		if r.voteCount > len(r.nodes)/2+1 {
			r.toLeaderC <- true
		}
	}
}

func (r *Raft) RequestVote(args VoteArgs, reply *VoteReply) error {

	if args.Term < r.currentTerm {
		reply.Term = r.currentTerm
		reply.VoteGranted = false
		return nil
	}

	if r.votedFor == -1 {
		r.currentTerm = args.Term
		r.votedFor = args.CandidateID
		reply.Term = r.currentTerm
		reply.VoteGranted = true
		return nil
	}

	reply.Term = r.currentTerm
	reply.VoteGranted = false
	return nil
}

type HeartbeatArgs struct {
	Term     int
	LeaderID int
}

type HeartbeatReply struct {
	Term int
}

func (r *Raft) broadcastHeartbeat() {
	for i := range r.nodes {
		args := HeartbeatArgs{
			Term:     r.currentTerm,
			LeaderID: r.id,
		}

		go func(i int, args HeartbeatArgs) {
			var reply HeartbeatReply
			r.sendHeartbeat(i, args, &reply)
		}(i, args)
	}
}

func (r *Raft) sendHeartbeat(serverID int, args HeartbeatArgs, reply *HeartbeatReply) {
	client, err := rpc.DialHTTP("tcp", r.nodes[serverID].address)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	client.Call("Raft.Heartbeat", args, reply)

	if reply.Term > r.currentTerm {
		r.currentTerm = reply.Term
		r.state = follower
		r.votedFor = -1
	}
}

// Heartbeat follower response
func (r *Raft) Heartbeat(args HeartbeatArgs, reply *HeartbeatReply) error {
	if args.Term < r.currentTerm {
		reply.Term = r.currentTerm
		return nil
	}

	if args.Term > r.currentTerm {
		r.currentTerm = args.Term
		r.state = follower
		r.votedFor = -1
	}

	reply.Term = r.currentTerm

	r.heartbeatC <- true

	return nil
}

func (r *Raft) rpc(port string) {
	rpc.Register(r)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen err:", err)
	}
	go http.Serve(l, nil)
}

func main() {

	// 参数定义和默认值
	port := flag.String("port", ":9091", "rpc listen port")
	cluster := flag.String("cluster", "127.0.0.1:9091", "comma sep")
	id := flag.Int("id", 1, "node ID")

	flag.Parse()
	clusters := strings.Split(*cluster, ",")
	ns := make(map[int]*node)
	for k, v := range clusters {
		ns[k] = newNode(v)
	}

	Raft := &Raft{}
	Raft.id = *id
	Raft.nodes = ns

	Raft.rpc(*port)
	Raft.start()

	select {}

}
