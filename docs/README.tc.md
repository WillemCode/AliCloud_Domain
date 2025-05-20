# 阿里雲域名管理工具
🔍 **一個快速管理多賬號阿里雲域名解析的小工具**
自動識別域名所屬賬號，支援添加、刪除、修改、查詢解析記錄，解決多賬號域名管理的混亂問題。
---
## ✨ 功能特性
- **多賬號支援**：同時配置多個阿里雲 AK/SK，自動匹配域名所屬賬號。
- **一鍵操作**：無需手動切換賬號，直接管理域名解析。
- **核心功能**：
  - ✅ 添加域名解析記錄（A/CNAME/MX/TXT 等）
  - ✅ 刪除域名解析記錄
  - ✅ 修改域名解析記錄
  - ✅ 查詢域名詳細信息
  - 🔍 模糊搜索域名信息 (www/example.com/www.example.com)
- **輕量高效**：Go 語言開發，命令行交互，快速響應。
---
## 🚀 快速開始
### 安裝方式
#### 1. 直接運行（需 Go 環境）
```bash
git clone https://github.com/WillemCode/AliCloud_Domain.git
cd AliCloud_Domain
go run cmd/main.go
```
#### 2. 下載二進制文件（推薦）

---
### 配置步驟
1. **添加阿里雲賬號**
   在程序首次運行時，在項目目錄中創建 `config.yaml` 按以下格式輸入阿里雲 AccessKey（AK）和 SecretKey（SK），支援添加多個賬號。
```yaml
aliyun_accounts:
  - name: "域名阿里雲-01"
    regionId : "cn-hangzhou"
    access_key: "·····················"
    access_secret: "···························"
  - name: "域名阿里雲-02"
    regionId : "cn-beijing"
    access_key: "·····················"
    access_secret: "···························"
  - name: "域名阿里雲-03"
  ······
```
2. **選擇操作**
   運行後會出現交互式菜單：
   ```text
   請選擇你要操作的內容... [type to search]: 
   > 1. 添加域名解析記錄
     2. 刪除域名解析記錄
     3. 修改域名解析記錄
     4. 查詢域名詳細信息
     5. 模糊搜索域名信息
   ```
---
## 📸 使用示例
### 添加解析記錄
!
### 刪除解析記錄
![刪除解析記錄截圖](./images/003.png)
### 修改解析記錄
![修改解析記錄截圖](./images/004.png)
### 查詢解析記錄
![查詢解析記錄截圖](./images/005.png)
### 模糊搜索域名
![模糊搜索截圖](./images/006.png)
---
## ⚙️ 技術棧
- **語言**: Go
- **阿里雲 SDK**: [Alibaba Cloud Go SDK](https://github.com/aliyun/alibaba-cloud-sdk-go)
- **命令行交互**: [Cobra](https://github.com/spf13/cobra) [Pterm](https://github.com/pterm/pterm)
---
## 貢獻與支持
- **Issues**：如在使用過程中遇到問題或有新需求，歡迎在倉庫的 Issue 中提出。
- **Pull Requests**：歡迎參與貢獻新的功能、優化或文檔修復。
---
## 📜 開源協議
本項目採用 [GNU General Public License (GPL)](./LICENSE) 進行開源發布。
這意味著：
- 你可以自由複製、修改和分發本項目的源代碼，但修改後的項目也必須繼續以 GPL 或兼容的許可證進行發布；
- 分發或發布時，需包含本項目的原始版權聲明與 GPL 協議文本，並提供完整的源代碼獲取方式。
---
## 🙋 常見問題
**Q: 如何保證 AK/SK 的安全性？**
A: 所有憑據僅存儲在本地配置文件中（如 `config.yaml`），**不會上傳到服務器**。
**Q: 支援國際版阿里雲嗎？**
A: 是，修改 SDK 的 `Endpoint` 即可。
---
## Star History
[![Star History Chart](https://api.star-history.com/svg?repos=WillemCode/AliCloud_Domain&type=Date)](https://www.star-history.com/#WillemCode/AliCloud_Domain&Date)
