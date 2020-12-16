document.querySelectorAll(".drop-zone__input").forEach(inputElement => {
    const dropZoneElement = inputElement.closest(".drop.zone");

    // TODO: fix error
    dropZoneElement.addEventListener("dragover", e => {
        dropZoneElement.classList.add("drop-zone--over");
    });
});