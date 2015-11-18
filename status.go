/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"regexp"
	"strings"
)

type Status struct {
	ClientsConnected    string
	ClientsBlocked      string
	MemoryUsed          string
	MemoryAllocated     string
	MemoryFragmentation string
	TotalCommands       string
	TotalExpired        string
	TotalEvicted        string
	Hits                string
	Misses              string
	Keys                string
	Expires             string
}

var ClientsConnectedRX, _ = regexp.Compile(`connected_clients:(\d+)`)
var ClientsBlockedRX, _ = regexp.Compile(`blocked_clients:(\d+)`)
var MemoryUsedRX, _ = regexp.Compile(`used_memory:(\d+)`)
var MemoryAllocatedRX, _ = regexp.Compile(`used_memory_rss:(\d+)`)
var MemoryFragmentationRX, _ = regexp.Compile(`mem_fragmentation_ratio:(\d+)`)
var TotalCommandsRX, _ = regexp.Compile(`total_commands_processed:(\d+)`)
var TotalExpiredRX, _ = regexp.Compile(`expired_keys:(\d+)`)
var TotalEvictedRX, _ = regexp.Compile(`evicted_keys:(\d+)`)
var HitsRX, _ = regexp.Compile(`keyspace_hits:(\d+)`)
var MissesRX, _ = regexp.Compile(`keyspace_misses:(\d+)`)
var KeysRX, _ = regexp.Compile(`db0:keys=(\d+),expires=(\d+).*`)

func parse(page string) Status {
	var result = Status{}

	var statusData = strings.Split(page, "\n")
	for _, line := range statusData {
		if ClientsConnectedRX.MatchString(line) {
			result.ClientsConnected = ClientsConnectedRX.FindStringSubmatch(line)[1]
		}
		if ClientsBlockedRX.MatchString(line) {
			result.ClientsBlocked = ClientsBlockedRX.FindStringSubmatch(line)[1]
		}

		if MemoryUsedRX.MatchString(line) {
			result.MemoryUsed = MemoryUsedRX.FindStringSubmatch(line)[1]
		}
		if MemoryAllocatedRX.MatchString(line) {
			result.MemoryAllocated = MemoryAllocatedRX.FindStringSubmatch(line)[1]
		}
		if MemoryFragmentationRX.MatchString(line) {
			result.MemoryFragmentation = MemoryFragmentationRX.FindStringSubmatch(line)[1]
		}

		if TotalCommandsRX.MatchString(line) {
			result.TotalCommands = TotalCommandsRX.FindStringSubmatch(line)[1]
		}
		if TotalExpiredRX.MatchString(line) {
			result.TotalExpired = TotalExpiredRX.FindStringSubmatch(line)[1]
		}
		if TotalEvictedRX.MatchString(line) {
			result.TotalEvicted = TotalEvictedRX.FindStringSubmatch(line)[1]
		}

		if HitsRX.MatchString(line) {
			result.Hits = HitsRX.FindStringSubmatch(line)[1]
		}
		if MissesRX.MatchString(line) {
			result.Misses = MissesRX.FindStringSubmatch(line)[1]
		}

		if KeysRX.MatchString(line) {
			result.Keys = KeysRX.FindStringSubmatch(line)[1]
			result.Expires = KeysRX.FindStringSubmatch(line)[2]
		}
	}

	return result
}
