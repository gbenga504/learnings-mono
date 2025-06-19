function rotateElementInArray(array, amountOfRotation) {
  for (
    let currentRotation = 1;
    currentRotation <= amountOfRotation;
    currentRotation++
  ) {
    for (let itemIndex = 0; itemIndex < array.length - 1; itemIndex++) {
      let temp = array[itemIndex];
      let previousIndex = itemIndex === 0 ? array.length - 1 : itemIndex - 1;

      array[itemIndex] = array[previousIndex];
      array[previousIndex] = temp;
    }
  }
}

// node rotateElementsInArray "[1, 2, 3, 4, 5, 6]" 2
// The result is ===> [ 3, 4, 5, 6, 1, 2 ]
function main() {
  try {
    const array = JSON.parse(process.argv[2]);
    const amountOfRotation = Number(process.argv[3]);

    rotateElementInArray(array, amountOfRotation);

    console.log("The result is ===>", array);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
