console.log("Release the hounds!!");


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
});