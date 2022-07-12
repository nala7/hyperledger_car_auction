from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class Car(BaseModel):
    id: str
    brand: str
    price: float
    color: str

class User(BaseModel):
    id: str
    name: str

# Crear
@app.post("/create")
def create_auction(car: Car, owner: User):
    return {"item_name": car.brand, "item_id": car.id, "owner_id": owner.id}

# Obtener
@app.get("/")
def get_cars():
    return {"items": []}

# Iniciar
@app.post("/start")
def start_auction(car: Car):
    return {"item_name": car.brand, "item_id": car.id}

# Cerrar
@app.post("/close")
def close_auction(car: Car):
    return {"item_name": car.brand, "item_id": car.id}

# Validar
@app.post("/validate")
def validate_auction(car: Car):
    return {"item_name": car.brand, "item_id": car.id}

# Pujar
@app.post("/bid")
def bid(car: Car, amount: int):
    return {"item_name": car.brand, "item_id": car.id, "amount": amount}

# Registrar usuario
@app.post("/register")
def register_user(user: User):
    return {"user_name": user.name}
