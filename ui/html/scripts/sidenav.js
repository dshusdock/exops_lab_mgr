export default () => ({ 
    list_search: "",
    chevronRotated: false,
    nodeListCopy: [],
    strLength: 0,
    test: "testing",

    tester() {
      console.log("tester");
    },

    onElementClick(el) {
      let children = el.childNodes;

      children.forEach((element) => {
        if (element.tagName === "I") {
          element.className =
            element.className === "fa fa-chevron-right rotate_back"
              ? "fa fa-chevron-right rotate_fwd"
              : "fa fa-chevron-right rotate_back";
        }
      });

      this.nodeListCopy = [];
    },
    onSearchChange(event) {
      let el = document.getElementsByClassName("list_container__ul")[0];
      const list = el.childNodes;

      // copy original list to nodeListCopy
      if (this.nodeListCopy.length === 0) {         
        list.forEach(function (currentValue, currentIndex, listObj) {           
          let copy = currentValue.cloneNode(true);
          this.nodeListCopy.push(copy);          
        }, this);
      }
      
      // purge the list
      while (el.firstChild) {
        el.removeChild(el.firstChild);
      }

      // filter the list
      this.nodeListCopy.forEach(function (currentValue, currentIndex, listObj) {
        if (currentValue.tagName === "LI") {                       
          if (currentValue.innerText.includes(this.list_search)) {              
            el.appendChild(currentValue);
          }
        } else {
          el.appendChild(currentValue);
        }
      }, this);        
    },    
})