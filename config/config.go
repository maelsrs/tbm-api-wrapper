package config

// gateway-apim.infotbm.com returns XML, while ws-*** returns JSON

const ( // Base
	TBM_BASE_URL    = "http://infotbm.com/"
	TBM_ALL_LINES   = "https://gateway-apim.infotbm.com/maas-web/web/v11/timetables/lines"
	TBM_MAJOR_EVENT = "https://gateway-apim.infotbm.com/maas-web/web/v11/traffic-info/major-event"
)

const (
	LINE_PRACTICAL_INFORMATION = "https://gateway-apim.infotbm.com/maas-web/web/v11/timetables/lines/[id]/practical-information"
	LINE_ROUTES                = "https://ws-maas.infotbm.com/ws/1.0/network/line-informations/[id]"

	ROUTE_TIMETABLE = "https://gateway-apim.infotbm.com/maas-web/web/v11/timetables/lines/[lineId]/routes/[routeId]/stops/[stopId]"
	// routeId is lineId or lineId_R (other way)

	STOP_TIMETABLE     = "https://gateway-apim.infotbm.com/maas-mobile/mobile/v11/timetables/stops/[stopId]"
	VEHICLE_NEXT_STOPS = "https://gateway-apim.infotbm.com/maas-web/web/v11/timetables/stops/[stopId]/next-stops-departures?vehicleJourneyId=[vehicleId]&departure=[departureTime]"
	// departureTime format: 2026-01-21T17:09:07+01:00 (2026-01-21T17%3A09%3A07%2B01%3A00)
	// vehicleId format: vehicle_journey:BMA:50227595-2026_HIVER-TRA_A410-Lun-Mer-30_dst_1
)

const (
	TRAFFIC_INFORMATION = "https://gateway-apim.infotbm.com/maas-web/web/v5/traffic-info?filter=[filter]"
	// filter values: TODAY, FUTURE
)
