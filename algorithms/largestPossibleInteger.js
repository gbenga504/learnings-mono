function largestPossibleInterger(array) {
  let result = "";
  // The maximum possible number can only be calculated if we go from 9 ===> 0
  // E.g First pass [9, 918, 10, 6]

  for (let i = 9; i >= 0; i--) {
    // const digitsThatMatches =  [9, 918]
    const digitsThatMatches = array.reduce((acc, num) => {
      const stringBasedNum = num.toString();

      if (Number(stringBasedNum.at(0)) === i) {
        acc.push(num);
      }

      return acc;
    }, []);

    // const lengthOfMaxDigit = 3
    const lengthOfMaxDigit = Math.max(...digitsThatMatches).toString().length;
    const dictOfDigitsAndAugumentation = {};

    // const augumentedDigitsThatMatches = [999, 918]
    // const dictOfDigitsAndAugumentation = {999: 9, 918: 918} ===> The key is the augumentation and the value is the actual digit
    const augumentedDigitsThatMatches = digitsThatMatches.map((d) => {
      const augumentedDigit = d.toString().padEnd(lengthOfMaxDigit, `${i}`);
      dictOfDigitsAndAugumentation[augumentedDigit] = d;

      return Number(augumentedDigit);
    });

    // result = "9918"
    // Here we get the max augumented digits and append to actual digit to result
    while (augumentedDigitsThatMatches.length > 0) {
      const maxDigit = Math.max(...augumentedDigitsThatMatches);
      const indexOfMaxDigit = augumentedDigitsThatMatches.indexOf(maxDigit);

      result += dictOfDigitsAndAugumentation[maxDigit].toString();
      augumentedDigitsThatMatches.splice(indexOfMaxDigit, 1);
    }
  }

  return result;
}

function main() {
  try {
    let array = JSON.parse(process.argv[2]).slice(0);
    const result = largestPossibleInterger(array);

    console.log(result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

// node largestPossibleInteger "[10, 7, 76, 415]"
// gives 77641510
main();
