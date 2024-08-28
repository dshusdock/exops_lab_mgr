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
      console.log("close clicked");
      let modal = document.getElementById("myModal");
      modal.style.display = "none";
    },
    onOutsideClick(event) {
      console.log("outside clicked");
      let modal = document.getElementById("myModal");
      if (event.target == modal) {
        modal.style.display = "none";
      }
    }
})