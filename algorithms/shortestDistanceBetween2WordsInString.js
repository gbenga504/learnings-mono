function shortestDistanceBetween2WordsInString({
  firstWord,
  secondWord,
  sentence,
}) {
  // Split the sentence into an array based on commas
  // Loop through the sentence array until you find the firstWord, skip otherwise
  // ===> Loop from the indexOfFirstWord + 1 ... end of array
  // ========> Find the second word or skip otherwise.
  // ============> Compute indexOfSecondWord - indexOfFirstWord - 1, then replace result if the answer is less than the current result
  // Finally return result
  // Explanation: Anytime we find the first word, we throw out some kind of beams to find the second word
  // And calculate the shortest distance

  let result = Infinity;
  const sentenceArray = sentence.split(/\s/);

  sentenceArray.forEach((word, index) => {
    if (word === firstWord && index !== sentenceArray.length - 1) {
      for (let j = index + 1; j <= sentenceArray.length - 1; j++) {
        if (sentenceArray[j] === secondWord) {
          const newResult = j - index - 1;

          if (newResult < result) {
            result = newResult;
          }
        }
      }
    }
  });

  return result;
}

function main() {
  try {
    const sentence = process.argv[2];
    const firstWord = process.argv[3];
    const secondWord = process.argv[4];

    const result = shortestDistanceBetween2WordsInString({
      firstWord,
      secondWord,
      sentence,
    });

    console.log("The result is ===>", result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
