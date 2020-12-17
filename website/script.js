const myForm = document.getElementById("myForm")
const inpFile = document.getElementById("inpFile")

myForm.addEventListener("submit", e => {
    e.preventDefault()

    const endpoint = 'http://localhost:8080/upload'
    const formData = new FormData()

    // inpFile come form name="inpFile"
    formData.append("inpFile", inpFile.files[0])

    fetch(endpoint, {
        method: "POST",
        body: formData,
        mode: 'no-cors',
    })
    .catch(console.error)
})
