#!/bin/bash

sed -i'' -e 's/json:"timeStamp"/json:"timeStamp" yaml:"-"/g' $1
sed -i'' -e 's/json:"commitHash"/json:"commitHash" yaml:"-"/g' $1
sed -i'' -e 's/json:"jiraIssue"/json:"jiraIssue" yaml:"-"/g' $1
sed -i'' -e 's/json:"owner"/json:"owner" yaml:"-"/g' $1
sed -i'' -e 's/json:"type,omitempty"/json:"type,omitempty" yaml:"-"/g' $1
