function minimumCostToBuildNHousesWithKColors(costs) {
  //If the matrix we are given only contains a single row, then it means that we are concerned with just one house
  //hence we find the smallest number in the array
  if (costs.length === 1) {
    return Math.min(...costs[0]);
  }

  //If the matrix has more rows then we try to reduce it to a single row
  while (costs.length > 1) {
    const currentRowCosts = costs[0];
    const nextRowCosts = costs[1];

    //We get the next row, remove the item in the currentIndex from the next row so we avoid cases where we
    //have computations with adjacent colors. Then we take the minimum and add to the current cost
    let newRowCosts = currentRowCosts.map((currentCost, currentCostIndex) => {
      const copiedNextRowCosts = [...nextRowCosts];
      copiedNextRowCosts.splice(currentCostIndex, 1);

      const minimumCost = Math.min(...copiedNextRowCosts);

      return currentCost + minimumCost;
    });

    //Delete the first 2 items from the matrix and add the new cost that has been computed
    costs.splice(0, 2);
    costs.unshift(newRowCosts);
  }

  return Math.min(...costs[0]);
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]);
    const result = minimumCostToBuildNHousesWithKColors(array);

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
