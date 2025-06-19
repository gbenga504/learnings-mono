function convertBinaryTreeToFullBinaryTree(root, parent = null) {
  // We have reached a leave node, hence we stop all operations i.e do nothing
  if (!root.left && !root.right) {
    return;
  }

  // If we have right node but not the left, then the node has 1 child and we detach it
  if (root.right && !root.left) {
    // Since in a binary tree, the left nodes take precedence over the right nodes especially for the inoder walk
    // we need to check if the parent has a left node, if YES, the we assign to the parent right node else we use the left node
    let assignedNode = null;

    if (!parent.left) {
      parent.left = root.right;
      assignedNode = "LEFT";
    } else {
      parent.right = root.right;
      assignedNode = "RIGHT";
    }

    convertBinaryTreeToFullBinaryTree(
      assignedNode === "LEFT" ? parent.left : parent.right,
      parent
    );
  }

  // If we have left node but not the right, then the node has 1 child and we detach it
  if (root.left && !root.right) {
    parent.left = root.left;

    convertBinaryTreeToFullBinaryTree(parent.left, parent);
  }

  // If we have both the left and right nodes, then we continue by iterating with them
  convertBinaryTreeToFullBinaryTree(root.left, root);
  convertBinaryTreeToFullBinaryTree(root.right, root);
}

function main() {
  try {
    let binaryTree = JSON.parse(process.argv[2]);
    convertBinaryTreeToFullBinaryTree(binaryTree);

    console.log(binaryTree);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
