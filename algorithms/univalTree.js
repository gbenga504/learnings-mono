function findNumberOfUnivalSubTree(root, result) {
  if (root.data === null) {
    return true;
  }

  const isLeftSubTreeAUnival = findNumberOfUnivalSubTree(root.left, result);
  const isRightSubTreeAUnival = findNumberOfUnivalSubTree(root.right, result);

  if (
    isLeftSubTreeAUnival &&
    isRightSubTreeAUnival &&
    root.left.data === null &&
    root.right.data === null
  ) {
    result.numberOfUnivalSubTree += 1;
    return result.isTreeUnival && true;
  }

  const areDataEqual =
    root.data === root.left.data && root.data === root.right.data;

  if (isLeftSubTreeAUnival && isRightSubTreeAUnival && areDataEqual) {
    result.numberOfUnivalSubTree += 1;
    return result.isTreeUnival && true;
  }

  return result.isTreeUnival && false;
}

function main() {
  try {
    let result = { numberOfUnivalSubTree: 0, isTreeUnival: true };
    const binaryTree = JSON.parse(process.argv[2]);

    findNumberOfUnivalSubTree(binaryTree, result);
    console.log("The result is ", result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
