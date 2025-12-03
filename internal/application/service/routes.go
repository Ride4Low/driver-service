package service

// {latitude: 13.747857208858562, longitude: 100.50850314787347}

// Predefined routes for drivers (used for the gRPC Streaming module)
// Center: {latitude: 13.747857208858562, longitude: 100.50850314787347}
// All routes are within 10 km radius
var PredefinedRoutes = [][][]float64{
	// Route 1: North-South route (3 waypoints)
	{
		{13.747857, 100.508503}, // Start
		{13.780234, 100.515678}, // ~3.7 km north
		{13.812456, 100.522145}, // ~7.2 km north
	},
	// Route 2: East-West route (8 waypoints)
	{
		{13.747857, 100.508503}, // Start
		{13.755234, 100.530145}, // ~2.4 km east
		{13.762145, 100.551234}, // ~4.7 km east
		{13.768923, 100.572345}, // ~7.0 km east
		{13.762145, 100.551234}, // Return
		{13.755234, 100.530145}, // Return
		{13.747857, 100.508503}, // Back to start
		{13.740234, 100.486789}, // ~2.4 km west
	},
	// Route 3: Circular route (8 waypoints)
	{
		{13.747857, 100.508503}, // Start
		{13.770234, 100.520145}, // NE
		{13.780456, 100.508503}, // N
		{13.770234, 100.496861}, // NW
		{13.747857, 100.484719}, // W
		{13.725480, 100.496861}, // SW
		{13.715258, 100.508503}, // S
		{13.725480, 100.520145}, // SE
	},
	// Route 4: Zigzag route (8 waypoints)
	{
		{13.747857, 100.508503}, // Start
		{13.758234, 100.515678}, // NE
		{13.737480, 100.530145}, // SE
		{13.768923, 100.535234}, // NE
		{13.758456, 100.545345}, // E
		{13.748123, 100.558789}, // E
		{13.738234, 100.548678}, // SE
		{13.728456, 100.538567}, // S
	},
}
