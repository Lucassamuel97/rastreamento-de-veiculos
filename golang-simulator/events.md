### Events Recebido

RouteCreated
- id
- distance
- directions []
    -- lat  
    -- lng


### Executar e retornar outro evento

FreightCalculated 
- route_id
- amount


--
DeliveryStarted
- route_id

### Evento retornado

DriverMoved
- route_id
- lat
- lng