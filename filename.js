var fs = require('fs');

let username = "test";
let dir = "./lambda";

// get files
fs.readdir(dir, function (err, files) {
    if (err) throw err;
    // iterate files
    files.forEach(function (file) {
        if (file.match(/username/)) {
            // create path
            var fileArray = file.split("-");
            fileArray[0] = username;
            let old_filepath = `${dir}/${file}`;
            let new_filepath = `${dir}/${fileArray.join("-")}`;
            // filename convert
            fs.rename(old_filepath, new_filepath, function (err) {
                if (err) {
                    console.error(err);
                    process.exit(1);
                }
                else {
                    console.log(old_filepath, new_filepath, 'finished!!');
                }
            });
        }
    })
});
