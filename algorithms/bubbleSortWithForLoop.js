function bubbleSort(array) {
  for (let i = 0; i <= array.length; i++) {
    for (let j = i; j <= array.length; j++) {
      if (array[j] < array[i] && array[j] !== array[i]) {
        let temp = array[i];

        array[i] = array[j];
        array[j] = temp;
      }
    }
  }
  return array;
}

function main() {
  try {
    let array = JSON.parse(process.args[2]).slice(0);
    bubbleSort(array);

    console.log(array);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}
