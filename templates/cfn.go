package templates

// Lambda template
// TODO overrides
// - Bucket location
// - custom policies
var Cfn = `
AWSTemplateFormatVersion: '2010-09-09'
Description: Golang Guardian lambda template

Parameters:
  Stage:
    Type: String
    Default: PROD
  Vpc:
    Type: AWS::EC2::VPC::Id
  Subnets:
    Type: List<AWS::EC2::Subnet::Id>
  Bucket:
    Type: String
  Key:
    Type: String
  Main:
    Type: String

Resources:
  Role:
    Type: AWS::IAM::Role
    Properties:
      ManagedPolicyArns:
        - arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - lambda.amazonaws.com
            Action:
              - sts:AssumeRole
      Path: /
      Policies:
        - PolicyName: LambdaPolicy
          PolicyDocument:
            Statement:
              - Effect: Allow
                Action:
                  - cloudwatch:*
                Resource: '*'
              - Effect: Allow
                Action:
                  - logs:CreateLogGroup
                  - logs:CreateLogStream
                  - logs:PutLogEvents
                Resource: '*'

  SecurityGroup:
    Type: "AWS::EC2::SecurityGroup"
    Properties:
      GroupDescription: "required as lambda must have one"
      VpcId: !Ref Vpc

  Lambda:
    Type: AWS::Lambda::Function
    Properties:
      Code:
        S3Bucket: !Ref Bucket
        S3Key: !Ref Key
      Description: Golang lambda
      Handler: !Ref Main
      MemorySize: 128
      Role: !GetAtt 'Role.Arn'
      Runtime: go1.x
      Timeout: 120
      VpcConfig:
        SecurityGroupIds:
          - !GetAtt SecurityGroup.GroupId
        SubnetIds: !Ref Subnets
`
