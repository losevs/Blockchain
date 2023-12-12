package server

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/losevs/Blockchain/chain"
)

type Message struct {
	BPM int `json:"bpm"`
}

func getBlockchain(c *fiber.Ctx) error {
	bytes, err := json.MarshalIndent(chain.Blockchain, "", "  ") //change prefix & suffix
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).Send(bytes)
}

func writeBlock(c *fiber.Ctx) error {
	var m Message
	if err := c.BodyParser(&m); err != nil {
		return respondJSON(c, fiber.StatusBadRequest, err)()
	}

	newBlock, err := chain.GenerateBlock(chain.Blockchain[len(chain.Blockchain)-1], m.BPM)
	if err != nil {
		return respondJSON(c, fiber.StatusInternalServerError, m)()
	}

	if chain.IsBlockValid(newBlock, chain.Blockchain[len(chain.Blockchain)-1]) {
		newBlockChain := append(chain.Blockchain, newBlock)
		chain.ReplaceChain(newBlockChain)
		spew.Dump(newBlock)
	}

	return respondJSON(c, fiber.StatusCreated, newBlock)()
}

func respondJSON(c *fiber.Ctx, code int, payload interface{}) func() error {
	return func() error {
		return c.Status(code).JSON(payload)
	}
}
