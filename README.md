
# Coiffeur Simulation

This project simulates a coiffeur (hairdresser) service system where clients arrive at random times, wait for an available server (coiffeur), and are served based on availability. The system uses a priority queue to handle the clients, assigning them to servers, processing their service, and tracking statistics like the total processing time, queue length, and number of clients served.

## Features

- Clients arrive at random times or at specified intervals.
- Clients are assigned to the available server with the least number of served clients.
- The system keeps track of the queue length and logs statistics like:
  - Total number of clients served.
  - Maximum queue length during the simulation.
  - Average processing time for all clients.

## Prerequisites

Before running the simulation, you will need to have the following:

- Go 1.18 or higher
- A terminal or command-line interface to run the Go program

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/RahmaKhlifi/coiffeur
   cd coiffeur
   ```

2. Make sure Go is installed on your system by running:

   ```bash
   go version
   ```

3. Install any necessary Go dependencies by running:

   ```bash
   go mod tidy
   ```

## Usage

### Running the Simulation with Specified Client Arrival Times

If you have a list of arrival times for the clients, you can use the `Run` method with a predefined list:

```go
sim := NewSimulator()
sim.AddServer("Coiffeur1", 1.5) // Server with 1.5 hours of processing time
sim.AddServer("Coiffeur2", 1.2) // Server with 1.2 hours of processing time

arrivalTimes := []float64{0.5, 1.0, 2.0, 3.0, 4.0} // List of client arrival times
sim.Run(arrivalTimes)
```

### Running the Simulation with Random Client Arrivals

Alternatively, you can simulate clients arriving at random times based on a rate parameter (`lambda`) and a duration. This is useful for performance testing with random input.

```go
sim := NewSimulator()
sim.AddServer("Coiffeur1", 1.5)
sim.AddServer("Coiffeur2", 1.2)

lambda := 1.0  // Average arrival rate (clients per time unit)
duration := 10.0  // Simulation duration
sim.RunWithRandomClients(lambda, duration)
```

### Logging

The simulation logs the following details:

- Client arrival time and assigned server.
- Client start and end processing times.
- Total processing time for each client.
- Queue length and maximum queue length during the simulation.
- Average processing time across all clients.

## Example Output

The output is logged to the console, and you should expect to see entries similar to the following:

```log
Client 1 arrived at time 0.50
Client 2 arrived at time 1.00
Client 3 arrived at time 2.00
Client 4 arrived at time 3.00
Client 5 arrived at time 4.00

Client 1 started processing at time 0.50 on Server Coiffeur1
Client 2 started processing at time 1.00 on Server Coiffeur2

... [Other client events] ...

Total clients served: 5
Maximum queue length during the simulation: 2
Average processing time: 1.50 hours
```


## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Create a new pull request.

