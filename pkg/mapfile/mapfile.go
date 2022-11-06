/*
Package mapfile contains utilities for loading/parsing .map files used in pathfinding benchmarks.
For more information on the .map format, please visit:

https://movingai.com/benchmarks/formats.html

---
.map file

	type octile -- always octile
	height y -- height of the map
	width x -- width of the map
	map -- the actual map data

The map data is store as an ASCII grid. The upper-left corner of the map is (0,0). The following characters are possible:

	. - passable terrain
	G - passable terrain
	@ - out of bounds
	O - out of bounds
	T - trees (unpassable)
	S - swamp (passable from regular terrain)
	W - water (traversable, but not passable from terrain)

---
.scen file

	version n.m -- n and m are integers. If m == 0, it can be ommitted
	[bucket] [map] [map-width] [map-height] [start x] [start y] [goal x] [goal y] [optimal length]

Some details:

	[optimal length]: assumes sqrt(2) diagonal costs, and that agents cannot cut corners through walls
	[map-width] [map-height]: if these do not match the width and height of the .map file, it should be scaled to that size.
	[map]: the path of a .map file

Example:

	version 1
	0	brc000d.map	257	261	75	193	77	195	2.82842712
	0	brc000d.map	257	261	126	193	126	194	1.00000000
	0	brc000d.map	257	261	8	36	8	38	2.00000000
	0	brc000d.map	257	261	253	163	251	163	2.00000000
	0	brc000d.map	257	261	170	79	169	76	3.41421356
*/
package mapfile
