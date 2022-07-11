package main

import(
  eventsProvider "client/providers"
)

func main() {
  em  := eventsProvider.EventsMiner{}
  em.Start()
}
