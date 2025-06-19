let totalPossiblePlacements = 0;

function nQueensProblem(n, rowId, columnIdsPlacedSoFar) {
  //This means we have move past the last row and have placed all the queens on the board hence
  //we can increase the total placement count and exit this function
  if (n === rowId) {
    return totalPossiblePlacements++;
  }

  //Since only one queen can be placed in a row or column. We attempt to place the first queen in the first column,
  //check for compatibility with other queens and iteratively move the first queen to second column repeating the process
  for (let columnId = 0; columnId < n; columnId++) {
    columnIdsPlacedSoFar.push(columnId);

    if (isPlacementvalid(columnIdsPlacedSoFar, rowId)) {
      //If the current placement is valid, then we go ahead to place another queen in the next row
      nQueensProblem(n, rowId + 1, columnIdsPlacedSoFar);
    }

    //Once we are done with validating the placement of the current queen, it is a good idea to
    //remove it from the array
    columnIdsPlacedSoFar.pop();
  }
}

function isPlacementvalid(columnIdsPlacedSoFar, rowId) {
  //Move thorough all former queen placements
  for (
    let columnIdIndex = 0;
    columnIdIndex < columnIdsPlacedSoFar.length - 1;
    columnIdIndex++
  ) {
    //check if the current placement is on the same column or diagonal to the previous queen
    //diaginal = rowIdOfCurrentPlacement - rowIdOfPreviousQueenPlacement === columnIdOfCurrentPlacement - columnIdOfPreviousQueenPlacement
    const diff = Math.abs(
      columnIdsPlacedSoFar[rowId] - columnIdsPlacedSoFar[columnIdIndex]
    );

    if (diff === 0 || diff === rowId - columnIdIndex) {
      return false;
    }
  }

  return true;
}

function main() {
  try {
    const n = Number(process.argv[2]);
    nQueensProblem(n, 0, []);

    console.log(totalPossiblePlacements);
  } catch (error) {
    console.log("Error occured:", error);
  }
}

main();
