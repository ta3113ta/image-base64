function dropHandler(e) {
    console.log('File(s) dropped');

    // Prevent default behavior (Prevent file from being opened)
    e.preventDefault();

    if (e.dataTransfer.items) {
        for (var i = 0; i < e.dataTransfer.items.length; i++) {
            // If dropped items aren't files, reject them
            if (e.dataTransfer[i].kind === 'file') {
                var file = e.dataTransfer.items[i].getAsFile();
                console.log('...file['+i+'].name = '+file.name);
            }
        }
    } else {
        // Use DataTransfer interface to access the file(s)
        for (var i = 0; i < e.dataTransfer.files.length; i++) {
            console.log('... file[' + i + '].name = ' + ev.dataTransfer.files[i].name);
        }
    }
}

function dragOverHandler(e) {
    console.log('File(s) in drop zone');

    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
}