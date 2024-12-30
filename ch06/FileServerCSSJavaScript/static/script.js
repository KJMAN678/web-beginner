function addText() {
    let child = document.createElement('li');
    child.append("New Item");

    let parent = document.getElementById('sample');
    parent.appendChild(child);
}
