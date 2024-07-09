console.log("Release the hounds!!......again!");

document.addEventListener("alpine:init", () => {
    console.log("modalData store initialized");
    Alpine.store("modalData", {
     
      btn: document.getElementById("myBtn"),
      span: document.getElementsByClassName("close")[0],     
      // Functions
      onFileClick(event) { 
        console.log("on file clicked");
        document.getElementById("fileInput").click();
      },
      onChange(event) {
        console.log("on change clicked" + event.target.value);
     
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
      
    }),
    console.log("hdrData store initialized");
    Alpine.store("hdrData", {
      
      btn: document.getElementById("myBtn"),
      span: document.getElementsByClassName("close")[0],     
      // Functions
      onClick(event) {
        console.log("hdr clicked");
        // let modal = document.getElementsByClassName("hdr__dropbtn")[0];
        let modal = document.getElementById("myDropdown");
        console.log("modal: ", modal);
        modal.style.display = "block";
      },
      onElementClick(event) {
        console.log("onElementClick clicked");
        let modal = document.getElementsByClassName("hdr__dropdown-content")[0];
        // let modal = document.getElementById("myDropdown");
        console.log("modal: ", modal);
        modal.style.display = "none";
      },
      onOutsideClick(event) {
        console.log("outside clicked");
        let modal = document.getElementById("myDropdown");
        modal.style.display = "none";
      }
      
    })
});
