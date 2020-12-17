document.querySelectorAll(".drop-zone__input").forEach(inputElement => {
    const dropZoneElement = inputElement.closest(".drop-zone");

    dropZoneElement.addEventListener("dragover", e => {
        e.preventDefault();
        dropZoneElement.classList.add("drop-zone--over");
    })


    let dragEvents = ['dragleave', 'dragend'];
    dragEvents.forEach(type => {
        dropZoneElement.addEventListener(type, e => {
            dropZoneElement.classList.remove("drop-zone--over");
        });
    });


    dropZoneElement.addEventListener("drop", e => {
        e.preventDefault();

        if (e.dataTransfer.files.length) {
            inputElement.files = e.dataTransfer.files;
            updateThumbnail(dropZoneElement, e.dataTransfer.files[0]);

            dropZoneElement.classList.remove("drop-zone--over");
        }
    });
});

function updateThumbnail(dropZoneElement, file) {
    let thumnailElement = dropZoneElement.querySelector(".drop-zone__thumb");

    // First time - there is no thumbnail element, so lets create it
    thumnailElement = document.createElement("div");
    thumnailElement.classList.add("drop-zone__thumb");
    dropZoneElement.appendChild(thumnailElement);
}