export default () => ({ 
    btn: document.getElementById("myBtn"),
    span: document.getElementsByClassName("close")[0],
    // Functions
    onClick(event) {
      let modal = document.getElementById("myDropdown");
      modal.style.display = "block";
    },
    onElementClick(event) {
      // Disappear the dropdown
      let modal = document.getElementsByClassName("hdr__dropdown-content")[0];
      modal.style.display = "none";
      // let sidnav = document.getElementsByClassName("sidenav")[0];
      // sidnav.style.display = "block";
    },
    onOutsideClick(event) {
      console.log("outside clicked");
      let modal = document.getElementById("myDropdown");
      modal.style.display = "none";
    }
})