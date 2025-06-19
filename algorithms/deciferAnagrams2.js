function doesWordExistInAnagram(word, anagram) {
  const result = word.split("").reduce((acc, letter) => {
    acc = acc && anagram.indexOf(letter) !== -1;

    return acc;
  }, true);

  return result;
}

function deciferAnagrams2(anagram, word) {
  let result = "";
  let anagramWordLengthToConsider = anagram.length - word.length;

  for (let i = 0; i <= anagramWordLengthToConsider; i++) {
    const anagramToCheck = anagram.substring(i, i + word.length);

    if (doesWordExistInAnagram(word, anagramToCheck)) {
      result += `${i}`;
    }
  }

  return result;
}

function main() {
  try {
    const anagrams = process.argv[2];
    const word = process.argv[3];
    const result = deciferAnagrams2(anagrams, word);

    console.log("The result is ===>", result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
