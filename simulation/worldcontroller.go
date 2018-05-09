package simulation

import (
	"github.com/soupstore/coda-world/simulation/model"
)

// WorldController is an interface over Simulation for modifying the world itself
type WorldController interface {
	CreateWorld(worldID model.WorldID) error
	DestroyWorld(worldID model.WorldID)
	CreateRoom(worldID model.WorldID, roomID model.RoomID, name, description string) (*model.Room, error)
	GetRoom(worldID model.WorldID, roomID model.RoomID) (*model.Room, error)
	DestroyRoom(worldID model.WorldID, roomID model.RoomID) error
	SetSpawnRoom(worldID model.WorldID, roomID model.RoomID) error
	CreateItem(itemID model.ItemID, name string, aliases []string, rigSlot model.RigSlot) (*model.Item, error)
	SpawnItem(itemID model.ItemID, containerID model.ContainerID) error
}

// CreateWorld creates a new world in the simulation.
// Every world must have a unique WorldID, which is a type aliased sting.
func (s *Simulation) CreateWorld(worldID model.WorldID) error {
	// TODO: check for uniqueness
	s.worlds[worldID] = model.NewWorld(worldID)
	return nil
}

// DestroyWorld unloads a world and all of its rooms from the simulation.
func (s *Simulation) DestroyWorld(worldID model.WorldID) {
	// TODO: Move all characters in this world to a safe location
	delete(s.worlds, worldID)
}

// CreateRoom creates a new room in the specified world with the specified room ID
func (s *Simulation) CreateRoom(worldID model.WorldID, roomID model.RoomID, name, description string) (*model.Room, error) {
	containerID := s.getNextContainerID()

	world, ok := s.worlds[worldID]
	if !ok {
		return nil, ErrWorldNotFound
	}

	// TODO: Check that room with ID does not already exist

	room := model.NewRoom(roomID, worldID, containerID, name, description)
	world.Rooms[roomID] = room

	container := room.Container
	s.containers[container.ID] = container

	return room, nil
}

// GetRoom returns the room object as the specified world ID and room ID
func (s *Simulation) GetRoom(worldID model.WorldID, roomID model.RoomID) (*model.Room, error) {
	world, ok := s.worlds[worldID]
	if !ok {
		return nil, ErrWorldNotFound
	}

	room, ok := world.Rooms[roomID]
	if !ok {
		return nil, ErrRoomNotFound
	}

	return room, nil
}

// DestroyRoom removes a room from a world.
func (s *Simulation) DestroyRoom(worldID model.WorldID, roomID model.RoomID) error {
	world, ok := s.worlds[worldID]
	if !ok {
		return ErrWorldNotFound
	}

	// TODO: Clean up broken exits in the rest of the sim.

	delete(world.Rooms, roomID)

	return nil
}

// SetSpawnRoom sets the room that all new characters will start in
func (s *Simulation) SetSpawnRoom(worldID model.WorldID, roomID model.RoomID) error {
	room, err := s.GetRoom(worldID, roomID)
	if err != nil {
		return ErrRoomNotFound
	}

	s.spawnRoom = room

	return nil
}

// CreateItem creates a new item archetype
func (s *Simulation) CreateItem(itemID model.ItemID, name string, aliases []string, rigSlot model.RigSlot) (*model.Item, error) {
	item := model.NewItem(itemID, name, aliases, rigSlot)
	s.items[itemID] = item
	return item, nil
}

// TODO: Clean up this, its messy
func (s *Simulation) SpawnItem(itemID model.ItemID, containerID model.ContainerID) error {
	container, ok := s.containers[containerID]
	if !ok {
		return ErrContainerNotFound
	}

	archetype := s.items[itemID]
	copy := *archetype

	container.PutItem(&copy)
	return nil
}
