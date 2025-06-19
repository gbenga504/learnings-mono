function doesWordExistInAnagram(word, anagram) {
  const result = word.split("").reduce((acc, letter) => {
    acc = acc && anagram.indexOf(letter) !== -1;

    return acc;
  }, true);

  return result;
}

function removeWordFromAnagram(word, anagram) {
  const result = word.split("").reduce((acc, letter) => {
    return acc.replace(letter, "");
  }, anagram);

  return result;
}

function deciferAnagrams(anagram) {
  let result = "";
  // create an object mapping containing numbers and their values written in words
  // Loop through the object; [2 for loops]
  // For each inner Loop, check if the number exist in the anagram.
  // While we can find the number in the anagram, we keep anagramming and removing found letters
  // after existing the while loop, repeat the process but with the next inner loop
  // If the number has not been recorded and result does not contain any recorded values then [break from inner loop]
  // In the bigger forloop, check if we deciphered, if we did, then break also and return result else continue
  const words = {
    0: "zero",
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
  };

  // We use 9 because the words are 9 in total
  for (let i = 0; i <= 9; i++) {
    let anagramCopy = anagram;
    result = "";

    for (let j = i; j <= 9; j++) {
      const word = words[j];

      while (doesWordExistInAnagram(word, anagramCopy)) {
        anagramCopy = removeWordFromAnagram(word, anagramCopy);
        result += `${j}`;
      }

      // We check if the word does not exist in the anagram
      // and the anagrams length has not changed, if so we break as soon as possible to avoid wasting time
      // Also if we have completely deciphered the anagram, we break
      if (anagramCopy === anagram || anagramCopy.length === 0) {
        break;
      }
    }

    if (anagramCopy.length === 0) {
      break;
    }
  }

  return result;
}

function main() {
  try {
    const anagrams = process.argv[2];
    const result = deciferAnagrams(anagrams);

    console.log("The result is ===>", result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
