function deepestNodeInBinaryTree(root) {
  if (!root.left && !root.right) {
    return { node: root, depth: 1 };
  }

  if (root.left && !root.right) {
    return incrementDepth(deepestNodeInBinaryTree(root.left));
  } else if (root.right && !root.left) {
    return incrementDepth(deepestNodeInBinaryTree(root.right));
  }

  return incrementDepth(
    maximumDepth(
      deepestNodeInBinaryTree(root.left),
      deepestNodeInBinaryTree(root.right)
    )
  );
}

function maximumDepth(leftTreeTuple, rightTreeTuple) {
  if (leftTreeTuple.depth > rightTreeTuple.depth) {
    return leftTreeTuple;
  }

  return rightTreeTuple;
}

function incrementDepth(tuple) {
  return { ...tuple, depth: tuple.depth + 1 };
}

function main() {
  try {
    const root = JSON.parse(process.argv[2]);
    const result = deepestNodeInBinaryTree(root);

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
