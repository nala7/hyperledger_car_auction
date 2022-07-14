from fastapi import FastAPI
from pydantic import BaseModel
import broker 

app = FastAPI()

class Car(BaseModel):
    id: str
    model: str
    colour: str

class User(BaseModel):
    id: str
    name: str

# Obtener todos los carros registrados
@app.get("/")
async def get_cars():
    cars = await broker.get_all_cars()
    return cars

# Obtener un carro por su id
@app.get("/cars/{car_id}")
async def get_cars(car_id: int):
    car = await broker.get_car(car_id)
    return car

### Usuarios vendedores ###

# Crear auto/subasta
@app.post("/create")
async def create_auction(car: Car, owner: User):
    new_car = await broker.create_car(car.id, car.model, car.colour, owner.id)
    return new_car

# Iniciar subasta de auto
@app.post("/start")
async def start_auction(car: Car, owner: User):
    auction = await broker.start_auction(car.id, owner.id)
    return auction

# Cerrar subasta
@app.post("/close")
async def close_auction(car: Car, owner: User):
    auction = await broker.close_auction(car.id, owner.id)
    return auction

# Validar
@app.post("/validate")
async def validate_auction(car: Car, owner: User):
    new_owner = await broker.verify_auction(car.id, owner.id)
    return new_owner

### Usuarios compradores ###
# Pujar
@app.post("/bid")
async def bid(car: Car, gambler: User, amount: int):    
    bid = await broker.bid(car.id, gambler.id, amount)
    return bid
