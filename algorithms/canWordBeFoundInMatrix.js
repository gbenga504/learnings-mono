function isTargetWordInRow(row, colIndex, targetWord) {
  // The colIndex indicates the particular letter in the row we want to search from
  const word = row.join("").substring(colIndex, colIndex + targetWord.length);

  return word.toLowerCase() === targetWord.toLowerCase();
}

function isTargetWordInColumn(array, colIndex, rowIndex, targetWord) {
  // The colIndex tells us how to form the word
  // The rowIndex Indicates the particular letter we want to search for
  let word = array.reduce((acc, row) => {
    acc += row[colIndex];

    return acc;
  }, "");

  word = word.substring(rowIndex, rowIndex + targetWord.length);

  return word.toLowerCase() === targetWord.toLowerCase();
}

function canWordBeFoundInMatrix(targetWord, array) {
  let foundTargetWord = false;

  for (let row = 0; row < array.length; row++) {
    for (let col = 0; col < array[row].length; col++) {
      const canSearchRow = array[row].length - col >= targetWord.length;
      const canSearchCol = array.length - row >= targetWord.length;

      if (!canSearchCol && !canSearchRow) {
        continue;
      }

      // We go ahead to search the row for the target word if possible
      if (canSearchRow && isTargetWordInRow(array[row], col, targetWord)) {
        foundTargetWord = true;

        break;
      }

      // We go ahead to search the col for the target word if possible
      if (canSearchCol && isTargetWordInColumn(array, col, row, targetWord)) {
        foundTargetWord = true;

        break;
      }
    }

    if (foundTargetWord) {
      break;
    }
  }

  return foundTargetWord;
}

function main() {
  try {
    const targetWord = process.argv[2];
    const array = JSON.parse(process.argv[3]);

    const result = canWordBeFoundInMatrix(targetWord, array);

    console.log(result);
  } catch (error) {
    console.log("Error occured:", error);
  }
}

main();
