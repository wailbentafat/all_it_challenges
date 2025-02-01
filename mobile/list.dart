Map<String, dynamic> calcule(List<int> list) {
  List<int> result = [];
  int sum = 0;

  for (int num in list) {
    if (num % 2 == 0) {
      result.add(num * num); 
    } else {
      sum += num; 
    }
  }

  return {"result": result, "sum": sum}; 
}

void main() {
  List<int> numbers = [1, 2, 3, 4, 5];
  var output = calcule(numbers);

  print(output["result"]); 
  print(output["sum"]);    
}
