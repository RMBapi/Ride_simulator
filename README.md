# Ride Backend Project (Ride Simulator)


This project implements version 0.0.1 of the Ride backend for Pathao. It includes REST APIs for managing riders and drivers, enabling riders to request rides and drivers to manage their availability.

## The available features are:

- Rider registration via phone number
- Driver registration via phone number
- Driver status (online/offline) capability
- Rider ride request functionality
- Driver ride end functionality
- Driver current status check
- User current status check


### Rider registration via phone number

#### request 

```
POST http://localhost:8080/api/v1/riders
content-type: application/json

{
	"phoneNumber" : "01611111117"	
}

```
#### response

```
status code : HTTP/1.1 201 Created

{
  "event": {
    "Id": 7,
    "PhoneNumber": "01611111117",
    "Type": "rider"
  },
  "message": "Rider Created"
}


```

### Driver registration via phone number

#### request 

```
POST http://localhost:8080/api/v1/drivers
content-type: application/json

{
	"phoneNumber" : "01834284411"      	
}

```
#### response

```
status code : HTTP/1.1 201 Created

{
  "event": {
    "Id": 5,
    "PhoneNumber": "01834284411",
    "Type": "driver"
  },
  "message": "Driver Created"
}

```


### Driver status (online/offline) capability


#### request (online)

```
PUT  http://localhost:8080/drivers/5/status
Content-Type: application/json


{
	"status" : "online"      
}

```
#### response (online)

```
status code : HTTP/1.1 200 OK

{
  "event": {
    "Status": "online",
    "Id": 5,
    "PhoneNumber": "01834284411",
    "Type": "driver"
  },
  "message": "Driver status updated"
}

```


#### request (offline)

```
PUT  http://localhost:8080/drivers/5/status
Content-Type: application/json


{
	"status" : "offline"      
}

```
#### response (offline)

```
status code : HTTP/1.1 200 OK

{
  "event": {
    "Status": "offline",
    "Id": 5,
    "PhoneNumber": "01834284411",
    "Type": "driver"
  },
  "message": "Driver status updated"
}

```

### Rider ride request functionality


#### request

```
POST http://localhost:8080/api/v1/riders
content-type: application/json

{
	"riderID" : 4
}


```
#### response 

```
status code : HTTP/1.1 201 Created

{
  "event": {
    "Id": 6,
    "RiderId": 4,
    "RiderPhoneNumber": "01711334271",
    "DriverId": 1,
    "DriverPhoneNumber": "01834284316",
    "Status": "start"
  },
  "message": "ride Created"
}


```

### Rider ride End functionality


#### request

```
POST http://localhost:8080/api/v1/riders
content-type: application/json

{
	"driverId" : 3
}


```
#### response 

```
status code : HTTP/1.1 201 Created

{
  "event": {
    "Id": 4,
    "RiderId": 3,
    "RiderPhoneNumber": "01711334284",
    "DriverId": 3,
    "DriverPhoneNumber": "01834284318",
    "Status": "end"
  },
  "message": "ride Ended"
}


```

### Driver current status check


#### request 

```
GET http://localhost:8080/api/v1/riders?driver_id=3

```
#### response (Driver online but not in the ride)

```
status code : HTTP/1.1 200 OK

{
  "message": "Driver Status",
  "status": "online"
}



```

#### response (Driver online and in a ride)

```
status code : HTTP/1.1 200 OK

{
  "event": {
    "Id": 0,
    "RiderId": 4,
    "RiderPhoneNumber": "01711334271",
    "DriverId": 1,
    "DriverPhoneNumber": "01834284316",
    "Status": "start"
  },
  "message": "Driver Status"
}



```

#### response (Driver offline)

```

{
  "message": "Driver Status",
  "status": "offline"
}

```


### User current status check


#### request

```
GET http://localhost:8080/api/v1/riders?rider_id=1

```
#### response (user in a ride)

```
status code : HTTP/1.1 200 OK

{
  "event": {
    "Id": 0,
    "RiderId": 1,
    "RiderPhoneNumber": "01711334282",
    "DriverId": 2,
    "DriverPhoneNumber": "01834284317",
    "Status": "start"
  },
  "message": "Rider Status"
}


```


#### response (user isn't a ride)

```
status code : HTTP/1.1 200 OK

{
  "message": "Rider didn't in a trip"
}


```







