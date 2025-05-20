# Alibaba Cloud Domain Management Tool

🔍 **A lightweight tool for managing multi-account Alibaba Cloud DNS records**

Automatically identifies domain ownership across accounts, supporting add/delete/modify/query operations to solve multi-account management chaos.

![UI Screenshot](./images/001.png)

---
## ✨ Features
- **Multi-account Support**: Configure multiple AK/SK pairs with automatic domain matching
- **One-click Operations**: Manage DNS records without manual account switching
- **Core Functions**:
  - ✅ Add DNS records (A/CNAME/MX/TXT etc.)
  - ✅ Delete DNS records
  - ✅ Modify DNS records
  - ✅ Query domain details
  - 🔍 Fuzzy domain search (www/example.com/www.example.com)
- **Lightweight & Efficient**: Built with Go, CLI interface, fast response
---
## 🚀 Quick Start
### Installation
#### 1. Run from source (requires Go)
```bash
git clone https://github.com/WillemCode/AliCloud_Domain.git
cd AliCloud_Domain
go run cmd/main.go
```
#### 2. Download binaries (recommended)
---
### Configuration
1. **Add Alibaba Cloud Accounts**
   Create `config.yaml` in project directory with this format:
```yaml
aliyun_accounts:
  - name: "Domain Account-01"
    regionId : "cn-hangzhou"
    access_key: "·····················"
    access_secret: "···························"
  - name: "Domain Account-02"
    regionId : "cn-beijing"
    access_key: "·····················"
    access_secret: "···························"
  - name: "Domain Account-03"
  ······
```
2. **Interactive Menu**
   After launch:
   ```text
   Select operation... [type to search]: 
   > 1. Add DNS record
     2. Delete DNS record
     3. Modify DNS record
     4. Query domain details
     5. Fuzzy domain search
   ```
---
## 📸 Examples
### Add Record
![Add Record](./images/002.png)
### Delete Record
![Delete Record](./images/003.png)
### Modify Record
![Modify Record](./images/004.png)
### Query Records
![Query Records](./images/005.png)
### Fuzzy Search
![Fuzzy Search](./images/006.png)
---
## ⚙️ Tech Stack
- **Language**: Go
- **Alibaba Cloud SDK**: [Alibaba Cloud Go SDK](https://github.com/aliyun/alibaba-cloud-sdk-go)
- **CLI**: [Cobra](https://github.com/spf13/cobra) [Pterm](https://github.com/pterm/pterm)
---
## Contribute & Support
- **Issues**: Report problems or request features via GitHub Issues
- **PRs**: Contributions welcome for new features/optimizations/docs
---
## 📜 License

Licensed under [GNU General Public License (GPL)](./LICENSE).
This means:
- You may copy, modify and distribute the software, but derivative works must remain under GPL
- Distributions must include original copyright notice and full source code

---
## 🙋 FAQ

**Q: How are AK/SK credentials secured?**

A: Stored only in local `config.yaml`, never uploaded.

**Q: International Alibaba Cloud supported?**

A: Yes, just modify the SDK's `Endpoint`.

---
## Star History
[![Star History Chart](https://api.star-history.com/svg?repos=WillemCode/AliCloud_Domain&type=Date)](https://www.star-history.com/#WillemCode/AliCloud_Domain&Date)
