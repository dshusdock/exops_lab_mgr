console.log("Release the hounds!!......again!");

document.addEventListener("alpine:init", () => {
    Alpine.store("myData", {
      target: "testing...",
      flag: true,
      drop: false,
      win1: true,
      win2: false,
      onClick() {
        console.log("clicked -- flag: ", this.flag);
        let el = document.getElementById("test_form");
        if (this.flag) {
          // el.className += " info-form--on"
          el.classList.add("info-form--on")
          this.flag = false;
        } else {
          el.classList.remove("info-form--on");
          el.classList.add("info-form--off");
          
          // el.className += " info-form--off"
          this.flag = true;
        }
      },
      handleTabClick(event) {
        console.log("clicked tab - " + event.target.innerText);

        let el = event.target;
        let clickedText = el.innerText;
        let parent = event.currentTarget;

        let nodes = parent.childNodes;
        console.log("nodes: ", nodes);

        nodes.forEach(element => {
          if (element.tagName === "SPAN") {
            console.log("element: ", element.innerText);

            if (clickedText === element.innerText) {
              console.log("found it: ", element);
              element.style.backgroundColor = "blue";
            } else { 
              console.log("found other: ", element.tagName);
              element.style.backgroundColor = "gray";
            }

            if (clickedText === "Test1") {
              this.win1 = true;
              this.win2 = false;
            } else {  
              this.win1 = false;
              this.win2 = true;
            }
          }
        });
      },
      onSettingClick(event) {
        this.drop = !this.drop;
      },
      testThis() {
        console.log("testThis: ", this);
        this.drop=false;
      },
      testThis2(event) {
        console.log("Got the focus");
        
      },
    })
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
      
    }),
    Alpine.store("hdrData", {
      
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
      
    }),
    Alpine.store("settings", {
      active: document.getElementsByClassName("page__general")[0],
      parm2: "",
      onMenuClick(event) {
        let text = event.target.innerText;
        console.log("menu clicked: ", text);

        if (text === "General") {
          console.log("General clicked");
          let modal = document.getElementsByClassName("page__general")[0];
          modal.style.display = "flex";  
          if(this.active) {
            this.active.style.display = "none";
            this.active = modal;        
          }
        }

        if (text === "Test") {
          console.log("Test clicked");
          let modal = document.getElementsByClassName("page__test")[0];
          modal.style.display = "flex";   
         
          this.active = document.getElementsByClassName("page__general")[0];
          if(this.active) {
            this.active.style.display = "none";
            this.active = modal;        
          }     
        }        
      },
    }),
    Alpine.store("tblehdr", {
      onHdrClick(event) {        
        // let modal = document.getElementById("myDropdown");
        console.log("header clicked" + event.target.innerText);
        console.log("header clicked-x" + event.clientX);
        console.log("header clicked-y" + event.clientY);
        let modal = document.getElementsByClassName("tbl-hdr-modal")[0];
        modal.style.display = "flex";
        modal.style.left = (event.clientX-250) + "px";
        modal.style.top = (event.clientY-100) + "px";
        let modalText = document.getElementsByClassName("tbl-hdr-modal__text")[0];
        modalText.innerText = event.target.innerText;
       
      },
      onCloseClick(event) {
        console.log("close clicked");
        let modal = document.getElementsByClassName("tbl-hdr-modal")[0];
        modal.style.display = "none";
      }


    }),
    Alpine.store("sidenav", {
        chevronRotated: false,
        onElementClick(el) {
            console.log("sidenav clicked" + el.innerText);
          
            let children = el.childNodes;
            console.log("children: ", children);
            children.forEach(element => {
              console.log("element: ", element.tagName);
              if (element.tagName === "I") {
                if (this.chevronRotated) {
                  element.className = "fa fa-chevron-right rotate_back";                  
                  this.chevronRotated = false;
                } else {
                  element.className = "fa fa-chevron-right rotate_fwd";                 
                  this.chevronRotated = true;
                }
                
                
              }
        })
      }


    })
  });
