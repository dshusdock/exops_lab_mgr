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
    })
  });
