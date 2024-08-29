export default () => ({ 
    onHdrClick(event) {
        let modal = document.getElementsByClassName("tbl-hdr-modal")[0];
        modal.style.display = "flex";
        modal.style.left = event.clientX - 250 + "px";
        modal.style.top = event.clientY - 100 + "px";
        let modalText = document.getElementsByClassName(
          "tbl-hdr-modal__text"
        )[0];
        modalText.innerText = event.target.innerText;
    },
    onCloseClick(event) {
        console.log("close clicked");
        let modal = document.getElementsByClassName("tbl-hdr-modal")[0];
        modal.style.display = "none";
    },    
})