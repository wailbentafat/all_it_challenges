class Car {
  String model;
  bool isAvailable;

  Car(this.model, {this.isAvailable = true});

  void rentCar() {
    if (isAvailable) {
      isAvailable = false;
      print("$model has been rented.");
    } else {
      print("$model is already rented.");
    }
  }

  void returnCar() {
    if (!isAvailable) {
      isAvailable = true;
      print("$model has been returned.");
    } else {
      print("$model was not rented.");
    }
  }
}

class RentalSystem {
  List<Car> inventory = [];

  void addCar(String model) {
    inventory.add(Car(model));
    print("$model has been added to the inventory.");
  }

  void rentCar(String model) {
    for (var car in inventory) {
      if (car.model == model && car.isAvailable) {
        car.rentCar();
        return;
      }
    }
    print("No available $model found.");
  }

  void returnCar(String model) {
    for (var car in inventory) {
      if (car.model == model && !car.isAvailable) {
        car.returnCar();
        return;
      }
    }
    print("No rented $model found.");
  }

  void displayCars() {
    print("\nCar Inventory:");
    for (var car in inventory) {
      String status = car.isAvailable ? "Available" : "Rented";
      print("- ${car.model}: $status");
    }
  }
}

void main() {
  RentalSystem rentalSystem = RentalSystem();


  rentalSystem.addCar("Toyota Corolla");
  rentalSystem.addCar("Honda Civic");
  rentalSystem.addCar("Ford Mustang");


  rentalSystem.displayCars();


  rentalSystem.rentCar("Honda Civic");

  rentalSystem.rentCar("Honda Civic");

  rentalSystem.returnCar("Honda Civic");


  rentalSystem.displayCars();
}
