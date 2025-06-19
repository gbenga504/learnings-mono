function maximumValueOfEachSubArray(array, numberToConsider) {
  if (numberToConsider === array.length) console.log([Math.max(...array)]);

  for (let i = 0; i < array.length - (numberToConsider - 1); i++) {
    const newArray = array.slice(i, i + numberToConsider);

    const largest = Math.max(...newArray);
    console.log(largest);
  }
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]);
    const numberToConsider = Number(process.argv[3]);

    maximumValueOfEachSubArray(array, numberToConsider);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
