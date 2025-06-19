function lengthOfLongestSubstring(word, numberOfDistinctCharacters) {
  let result = "";

  if (word.length === 1 && numberOfDistinctCharacters === 1) {
    return word;
  }

  let leftPointer = 0;
  let rightPointer = 0;
  let wordsSeenSoFar = "";
  let countOfDistinctLettersSeenSoFar = 0;

  function reset() {
    leftPointer++;
    rightPointer = leftPointer;
    countOfDistinctLettersSeenSoFar = 0;

    if (wordsSeenSoFar.length > result.length) {
      result = wordsSeenSoFar;
    }

    wordsSeenSoFar = "";
  }

  while (leftPointer <= word.length - 1) {
    if (
      countOfDistinctLettersSeenSoFar < numberOfDistinctCharacters &&
      wordsSeenSoFar.indexOf(word[rightPointer]) === -1
    ) {
      wordsSeenSoFar += word[rightPointer];
      countOfDistinctLettersSeenSoFar++;
      rightPointer + 1 <= word.length - 1 ? rightPointer++ : reset();
    } else if (
      countOfDistinctLettersSeenSoFar <= numberOfDistinctCharacters &&
      wordsSeenSoFar.indexOf(word[rightPointer]) !== -1
    ) {
      wordsSeenSoFar += word[rightPointer];
      rightPointer + 1 <= word.length - 1 ? rightPointer++ : reset();
    } else {
      reset();
    }
  }

  return result.length;
}

function usingSlidingWindow(word, numberOfDistinctCharacters) {
  let leftPointer = 0;
  let rightPointer = 0;
  let seen = new Map();
  let result = 0;

  while (rightPointer < word.length) {
    // If we have seen the word previously then we decrease the number of distinct characters
    if (!seen.has(word[rightPointer])) numberOfDistinctCharacters--;

    seen.set(word[rightPointer], rightPointer);

    //Once the numberOfDistict characters equal -1, then we need to run this
    while (numberOfDistinctCharacters < 0) {
      //This means we just saw a word we can remove from our dictionary to obtain balanace again
      if (leftPointer === seen.get(word[leftPointer])) {
        seen.delete(word[leftPointer]);
        numberOfDistinctCharacters++;
      }

      leftPointer++;
    }

    result = Math.max(result, rightPointer - leftPointer + 1);
    rightPointer++;
  }

  return result;
}

function main() {
  try {
    const word = process.argv[2];
    const numberOfDistinctCharacters = Number(process.argv[3]);
    const result = usingSlidingWindow(word, numberOfDistinctCharacters);

    console.log(result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
