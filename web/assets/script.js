let items = document.getElementsByTagName("li")

window.onload = function() {
    document.getElementById("input").focus();
}
// document.getElementById("inputfield").focus()
for (let i = 0; i < items.length; i++) {
    items[i].addEventListener("click", () => {
        items[i].classList.toggle("done");

        // Make a fetch request to the Golang server to toggle the Done value
        fetch('/toggle', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ index: i }),
        })
            .then(response => response.json())
            .then(data => {
                // Handle the updated list of todos if needed
                console.log(data);
            })
            .catch(error => console.error('Error:', error));
    });
 }