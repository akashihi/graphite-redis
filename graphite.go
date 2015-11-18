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
	"fmt"
	"github.com/marpaia/graphite-golang"
)

func sendMetrics(status Status, config Configuration) {
	var Graphite, err = graphite.NewGraphite(config.MetricsHost, config.MetricsPort)
	if err != nil {
		log.Error("Can't connect to graphite collector: %v", err)
		return
	}

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.clients.connected"), status.ClientsConnected)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.clients.blocked"), status.ClientsBlocked)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.memory.used"), status.MemoryUsed)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.memory.allocated"), status.MemoryAllocated)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.memory.fragmentation"), status.MemoryFragmentation)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.total.commands"), status.TotalCommands)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.total.expirations"), status.TotalExpired)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.total.evictions"), status.TotalEvicted)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.total.hits"), status.Hits)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.total.misses"), status.Misses)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.db0.keys"), status.Keys)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".redis.db0.expirableKeys"), status.Expires)
}
