// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package v1

type ConfigUpdateRequest struct {
	SiteCompanyAddress string `form:"site_company_address" json:"site_company_address"`
	SiteCompanyName    string `form:"site_company_name" json:"site_company_name"`
	SiteContactEmail   string `form:"site_contact_email" json:"site_contact_email"`
	SiteContactPhone   string `form:"site_contact_phone" json:"site_contact_phone"`
	SiteContactQq      string `form:"site_contact_qq" json:"site_contact_qq"`
	SiteContactWechat  string `form:"site_contact_wechat" json:"site_contact_wechat"`
	SiteDesc           string `form:"site_desc" json:"site_desc"`
	SiteIcp            string `form:"site_icp" json:"site_icp"`
	SiteKeywords       string `form:"site_keywords" json:"site_keywords" validate:"required"`
	SiteLogo           string `form:"site_logo" json:"site_logo"`
	SiteName           string `form:"site_name" json:"site_name" validate:"required"`
	SiteStatus         string `form:"site_status" json:"site_status"`
	SiteTitle          string `form:"site_title" json:"site_title" validate:"required"`
}
