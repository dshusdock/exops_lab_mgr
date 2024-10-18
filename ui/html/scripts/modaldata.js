export default () => ({ 
    btn: document.getElementById("myBtn"),
    span: document.getElementsByClassName("close")[0],
    // Functions
    onFileClick(event) {
      console.log("on file clicked");
      document.getElementById("fileInput").click();
    },
    onChange(event) {
      console.log("on change clicked" + event.target.value);
      let value = event.target.value;
      let fileName = value.split("\\").pop();
      console.log("fileName: ", fileName);

      document.getElementById("fileChoice").innerText = fileName;
    },
    onCloseClick(event) {
      console.log("modal close clicked");
      let modal = document.getElementById("myUploadModal");
      let x = document.getElementsByClassName("fileupload-modal")[0]
      x.style.display = "none";
      x.innerText = "TEST";
      modal.innerText = "TEST";
      modal.style.display = "none";
    },
    onOutsideClick(event) {
      console.log("outside clicked");
      let modal = document.getElementById("myUploadModal");
      if (event.target == modal) {
        modal.style.display = "none";
      }
    }
})