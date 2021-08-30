job "cw-redimail" {
  datacenters = ["pool-01"]

  group "website" {
    count = 5

    network {
      port "http" {
        static       = 8100
        to           = 8080
        host_network = "gateway"
      } 
    }
    
    task "server" {
      driver = "docker"

      config {
        image = "coldwireorg/redimail:v1.0.0"
        ports = ["http"]
        force_pull = true
      }
    }
  }
}
