package main

import(
  eventsProvider "client/events/providers"
)

func main() {
  em  := eventsProvider.EventsMiner{}
  em.Start()
}
