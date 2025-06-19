function countNumberOfNodesInABinaryTree(root) {
  if (root.data === null) return 0;

  let left = root.left ? countNumberOfNodesInABinaryTree(root.left) : 0;
  let right = root.right ? countNumberOfNodesInABinaryTree(root.right) : 0;

  return left + right + 1;
}

function main() {
  try {
    const root = JSON.parse(process.argv[2]);
    const numberOfNodes = countNumberOfNodesInABinaryTree(root);

    console.log(numberOfNodes);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
