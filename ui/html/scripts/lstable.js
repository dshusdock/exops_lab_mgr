export default () => ({ 
    someVar: "",
    info: { 
        name: "", 
        vip: "", 
        ip: "", 
        enterprise: "",
        role: "", 
        action: "" 
    },
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
    async onRowClick(el) {
        console.log("IN row clicked: ", el);
        const formData = new FormData(); 
        const myHeaders = new Headers();       

        let infoBox = document.getElementsByClassName("table-row-summary")[0];
        infoBox.classList.add("table-row-summary__on");
       
        let children = el.target.parentNode.childNodes;
        children.forEach((element, i) => {
            console.log("element: ", element.innerText + " - " + i);
            switch (i) {
            case 7:
                this.info.name = element.innerText;
                console.log("name: -->", this.info.name);
                document.getElementById("info_box1").innerText = this.info.name;
                break;
            case 11:
                this.info.ip = element.innerText;
                document.getElementById("info_box3").innerText = this.info.ip;
                break;
            case 13:
                this.info.vip = element.innerText;
                document.getElementById("info_box2").innerText = this.info.vip;
                break;
            case 21:
                this.info.enterprise = element.innerText;
                document.getElementById("info_box5").innerText = this.info.enterprise;
                break;
            case 23:
                this.info.role = element.innerText;
                document.getElementById("info_box6").innerText = this.info.role;
                break;            
            case 25:
                this.info.action = element.innerText;
                document.getElementById("info_box4").innerText = this.info.action;
                break;
            }          
        })

        myHeaders.append("Content-Type", "application/x-www-form-urlencoded");        
        formData.append("view_id", "unigystatus");
        formData.append("type", "request");
        formData.append("target", "ip");
        formData.append("data", this.info.ip);
        
        const myRequest = new Request("/request/status", {
            method: "POST",
            headers: myHeaders,
            body: new URLSearchParams(formData)
        });

        const response = await fetch(myRequest);
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }

        // console.log("response: ", response);

        const json = await response.json();
        console.log(json);
        if (json.Server === "RUNNING") {
            document.getElementById("thisnameel").style.color = "green";
        } else {
            document.getElementById("thisnameel").style.color = "red";
        }
    },
    onUMSClick() {
        console.log("UMS clicked");
        var strWindowFeatures = "location=yes,height=570,width=520,scrollbars=yes,status=yes";
        // var URL = "https://www.linkedin.com/cws/share?mini=true&amp;url=" + location.href;
        let URL = `https://${this.info.vip}/ums2/index.html?UMSClient=`
        var win = window.open(URL, "_blank");                
    },
    onDetailsClick() {
        console.log("Details clicked");
        var strWindowFeatures = "location=yes,height=570,width=520,scrollbars=yes,status=yes";
        // var URL = "https://www.linkedin.com/cws/share?mini=true&amp;url=" + location.href;
        let URL = `https://${this.info.vip}/haservices/checkHAStatus`
        var win = window.open(URL, "_blank");
    },
    onInfoLSideClick(ev) {
        console.log("info clicked" + ev.target.innerText);
        // navigator.clipboard.writeText(ev.target.innerText);
        // alert("Copied the text: " + ev.target.innerText);
        // Window.navigator.clipboard.writeText(copyText).then(function() {
        //     alert("Text copied to clipboard: " + copyText);
        // }).catch(function(error) {
        //     alert("Failed to copy text: " + error);
        // });        
    },
})