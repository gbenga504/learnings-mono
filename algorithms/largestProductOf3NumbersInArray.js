function largestProductOf3NumbersInArray(array) {
  if (array.length === 3) {
    return array[0] * array[1] * array[2];
  }

  let result = 0;

  for (let i = 0; i < array.length; i++) {
    for (let j = i + 1; j < array.length; j++) {
      for (let k = j + 1; k < array.length; k++) {
        const product = array[i] * array[j] * array[k];

        if (product > result) {
          result = product;
        }
      }
    }
  }

  return result;
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]);
    const result = largestProductOf3NumbersInArray(array);

    console.log(result);
  } catch (error) {
    console.log("Error occured:", error);
  }
}

main();
