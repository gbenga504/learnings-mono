class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
  }
}

class LinkedListForOrders {
  constructor(lastNOrders) {
    this.head = null;
    this.tail = null;
    this.lastNOrders = lastNOrders;
    this.currentOrders = 0;
  }

  record(order_id) {
    if (!this.head) {
      this.head = new Node(order_id);
      this.tail = this.head;
    } else {
      let currentNode = this.head;

      while (currentNode.next) {
        currentNode = currentNode.next;
      }

      currentNode.next = new Node(order_id);
      this.tail = currentNode.next;
    }

    // Increment count of currentOrders and if current orders is adding is going to be more than the
    // last n orders, then we want to delete the head node by re-ordering the pointer
    if (this.currentOrders + 1 <= this.lastNOrders) {
      this.currentOrders++;
    } else {
      this.head = this.head.next;
    }
  }

  getLast(i) {
    const indexOfNodesToSelect = this.lastNOrders - i;

    let count = 0;
    let selectedNode = this.head;

    while (count < indexOfNodesToSelect) {
      selectedNode = selectedNode.next;
      count++;
    }

    return { data: selectedNode.data };
  }
}

function main() {
  const myOrders = new LinkedListForOrders(5);

  myOrders.record(1);
  myOrders.record(2);
  myOrders.record(3);
  myOrders.record(4);
  myOrders.record(5);

  console.log(myOrders.getLast(5));

  myOrders.record(6);
  myOrders.record(7);

  console.log(myOrders.getLast(1));
}

main();
