package coiffeur

import (
	"log"
)

// ============
//  ClientQueue
// ============

type ClientQueue struct {
	items []*ClientInfo
}

func NewClientQueue() *ClientQueue {
	return &ClientQueue{items: []*ClientInfo{}}
}

func (q *ClientQueue) Push(client *ClientInfo) {
	q.items = append(q.items, client)
}

func (q *ClientQueue) Pop() *ClientInfo {
	if len(q.items) == 0 {
		return nil
	}
	var earliestClient *ClientInfo
	var index int
	for i, client := range q.items {
		if earliestClient == nil || client.arrivalTime < earliestClient.arrivalTime {
			earliestClient = client
			index = i
		}
	}
	q.items = append(q.items[:index], q.items[index+1:]...)
	return earliestClient
}

func (q *ClientQueue) Len() int {
	return len(q.items) + 1
}

// ============
//  Statistics
// ============

type ClientInfo struct {
	id                  int
	arrivalTime         float64
	startProcessingTime float64
	departureTime       float64
}

type Statistics struct {
	clientsInfo    []ClientInfo
	maxQueueLength int
}

// ========
//  Events
// ========

type ArrivalEvent struct {
	clientId    int
	arrivalTime float64
}

type StartProcessingEvent struct {
	clientId            int
	startProcessingTime float64
}

type EndProcessingEvent struct {
	serverId      string
	clientId      int
	departureTime float64
}

// =========
//  Server
// ========

type ServerStatus int

const (
	IDLE ServerStatus = iota // inactif
	BUSY                     // actif
)

type Server struct {
	sim            *Simulator
	id             string
	status         ServerStatus
	processingTime float64 // duration to achieve the service
	servedClients  int
}

func NewServer(sim *Simulator, id string, processingTime float64) *Server {
	return &Server{sim, id, IDLE, processingTime, 0}

}
func (this *Server) startProcessing(ev StartProcessingEvent) {
	if this.status != IDLE {
		log.Fatal("Server Must be IDLE !")
	}
	this.status = BUSY
	departureTime := ev.startProcessingTime + this.processingTime
	this.sim.InjectEvent(EndProcessingEvent{this.id, ev.clientId, departureTime})
}
func (this *Server) EndProcessing(ev EndProcessingEvent) {
	if this.status != BUSY {
		log.Fatal("Server Must be BUSY !")
	}
	this.servedClients++
	this.status = IDLE
}

// ===========
//  Simulator
// ===========

type Simulator struct {
	time        int
	servers     map[string]*Server // key = server ID
	queue       *EventQueue
	stats       Statistics
	clientQueue *ClientQueue // Use the ClientQueue instead of EventQueue
}

func NewSimulator() *Simulator {
	return &Simulator{
		0,
		make(map[string]*Server),
		NewEventQueue(),
		Statistics{clientsInfo: []ClientInfo{}, maxQueueLength: 0},
		NewClientQueue(),
	}
}

func (sim *Simulator) GetTime() int {
	return sim.time
}

func (sim *Simulator) AddServer(id string, processingTime float64) {
	server := NewServer(sim, id, processingTime)
	sim.servers[id] = server
}

// Add an event in the system
func (sim *Simulator) InjectEvent(event interface{}) {
	//	sim.queue..
	switch e := event.(type) {
	case ArrivalEvent:
		sim.queue.Push(&Event{event: e, priority: float64(e.arrivalTime)})
	case StartProcessingEvent:
		sim.queue.Push(&Event{event: e, priority: float64(e.startProcessingTime)})
	case EndProcessingEvent:
		sim.queue.Push(&Event{event: e, priority: float64(e.departureTime)})
	default:
		log.Fatal("unknown")
	}
}

func (sim *Simulator) ProcessEvent(event interface{}) {
	switch e := event.(type) {
	case ArrivalEvent:
		clientInfo := ClientInfo{
			id:          e.clientId,
			arrivalTime: e.arrivalTime,
		}
		sim.stats.clientsInfo = append(sim.stats.clientsInfo, clientInfo)

		log.Printf("Client %d arrived at time %.2f", e.clientId, e.arrivalTime)

		var availableServers []*Server = make([]*Server, 0)
		// Check for available servers
		for _, server := range sim.servers {
			if server.status == IDLE {
				availableServers = append(availableServers, server)
			}
		}

		// If a server is available, assign the client to it
		if len(availableServers) > 0 {
			var selectedServer *Server
			if len(availableServers) == 1 {
				selectedServer = availableServers[0]
			} else {
				// Select the server with the least number of served clients
				selectedServer = availableServers[0]
				for _, server := range availableServers {
					if server.servedClients < selectedServer.servedClients {
						selectedServer = server
					}
				}
			}

			// Log and start processing
			log.Printf("Client %d assigned to  %s for processing", e.clientId, selectedServer.id)

			startEvent := StartProcessingEvent{
				clientId:            e.clientId,
				startProcessingTime: e.arrivalTime,
			}
			selectedServer.startProcessing(startEvent)
			return
		}

		// No server available, client waits in the queue
		queueLength := sim.clientQueue.Len()
		if queueLength > sim.stats.maxQueueLength {
			sim.stats.maxQueueLength = queueLength
		}

		// Log the queue status
		log.Printf("No available Coiffeurs. Client %d is added to the queue (Queue Length: %d)", e.clientId, queueLength)
		// Add the client to the queue
		ci := &ClientInfo{
			id:          e.clientId,
			arrivalTime: e.arrivalTime,
		}
		sim.clientQueue.Push(ci)

	case StartProcessingEvent:
		// Start processing a client on a server
		for _, server := range sim.servers {
			if server.status == IDLE {
				server.startProcessing(e)
				log.Printf("Client %d started processing at time %.2f on  %s", e.clientId, e.startProcessingTime, server.id)

				// Update the statistics with the processing start time
				for i, client := range sim.stats.clientsInfo {
					if client.id == e.clientId {
						sim.stats.clientsInfo[i].startProcessingTime = e.startProcessingTime
						break
					}
				}
				return
			}
		}

	case EndProcessingEvent:
		// Handle end of processing
		for _, server := range sim.servers {
			if server.status == BUSY && server.id == e.serverId {
				server.EndProcessing(e)
				log.Printf("Client %d finished processing at time %.2f on  %s", e.clientId, e.departureTime, server.id)

				// Calculate the processing time
				var processingTime float64
				for _, client := range sim.stats.clientsInfo {
					if client.id == e.clientId {
						processingTime = e.departureTime - client.arrivalTime
						log.Printf("Client %d total processing time: %.2f hours", e.clientId, processingTime)
						break
					}
				}

				// Update the statistics with the departure time
				for i, client := range sim.stats.clientsInfo {
					if client.id == e.clientId {
						sim.stats.clientsInfo[i].departureTime = e.departureTime
						break
					}
				}

				// After processing, check if there are any clients waiting in the queue
				if sim.clientQueue.Len() > 0 {
					nextClient := sim.clientQueue.Pop() // Pop the next client from the queue
					if nextClient != nil {
						// Assign the next client to the server
						startEvent := StartProcessingEvent{
							clientId:            nextClient.id,
							startProcessingTime: e.departureTime, // Start time is after the previous client's processing
						}
						server.startProcessing(startEvent)

						// Log the assignment of the next client to the server
						log.Printf("Client %d is assigned to Server %s after the previous client finished", nextClient.id, server.id)
					}
				}

				return
			}
		}

	default:
		log.Fatalf("Unknown event type: %T", event)
	}
}

func (sim *Simulator) Run(ListeArriveClient []float64) {
	for i, tempsArrivee := range ListeArriveClient {
		event := ArrivalEvent{
			clientId:    i + 1,
			arrivalTime: tempsArrivee,
		}
		sim.InjectEvent(event)
	}

	for sim.queue.Len() > 0 {
		ev := sim.queue.Pop().event
		sim.ProcessEvent(ev)
	}

	sim.logStatistics()
}

func (sim *Simulator) logStatistics() {
	// Log the total number of served clients
	totalClients := len(sim.stats.clientsInfo)
	log.Printf("Total clients served: %d", totalClients)

	// Log the maximum queue length during the simulation
	log.Printf("Maximum queue length during the simulation: %d", sim.stats.maxQueueLength)

	// Calculate the average processing time for all clients
	var totalProcessingTime float64
	for _, client := range sim.stats.clientsInfo {
		totalProcessingTime += client.departureTime - client.arrivalTime
	}

	// Handle case where no clients were processed to avoid division by zero
	var averageProcessingTime float64
	if totalClients > 0 {
		averageProcessingTime = totalProcessingTime / float64(totalClients)
	}

	log.Printf("Average processing time: %.2f hours", averageProcessingTime)
}
