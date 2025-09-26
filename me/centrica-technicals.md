# Azure Batch-Based Risk Management System

## System Overview

The system, known as "Secure Environment," was built primarily on **Azure Batch** - a Platform-as-a-Service offering designed for coordinating massively scalable jobs, similar to what might be used today for AI training runs. The system orchestrated large-scale financial risk calculations and stress testing for what appears to be an energy trading company.

## Architecture Components

### Core Infrastructure
- **Azure Batch Coordinator**: Managed job distribution across virtual machines
- **Elastic Scaling**: Automatically spun up VMs to handle workload demands, maintaining a baseline of 4 VMs with ability to scale to thousands
- **Dependency Management**: Built-in directed acyclic graph (DAG) system for job dependencies
- **Dual Active Regions**: Both staging and production environments ran across two Azure regions for redundancy

### Job Processing Workflow
1. **Data Ingestion**: System pulled data from multiple sources including:
   - Foreign exchange rates
   - Energy prices
   - Wind turbine and power station meter readings
   - LNG vessel operational data
   - Trading data (hedged trades, options)
   - Financial feeds and futures information

2. **Data Preparation**: 
   - Reformatted raw data into required shapes for Python models
   - Fixed data quality issues and patched gaps
   - Packaged everything into zip files for deployment

3. **Job Execution**:
   - Deployed packages to shared storage accessible by Azure data centers
   - Azure Batch retrieved and executed Python entry points
   - Ran high-intensity mathematical calculations including Monte Carlo simulations

## Team Structure

**Quantitative Analysts Team**:
- Developed Python libraries for mathematical functions and models
- Created dependency graphs for parallel computations
- Defined data formatting requirements

**System Development Team**:
- Built the orchestration layer and scheduler (similar to CI/CD build runners)
- Managed data pipeline and preparation processes
- Designed stress testing system with user interface

## Key Use Cases

### Overnight Risk Calculations
- **Runtime**: Typically completed by 6 AM
- **Process**: Large-scale scenario testing and stress testing
- **Monitoring**: Required overnight supervision with gradual scale-down as jobs completed

### Stress Testing System
- Custom-designed interface for stress testing teams
- Scenario-based testing (e.g., tsunami impact on energy portfolio)
- Business user-friendly interface for defining test parameters

### Front Office Trading Support
- **Excel Integration**: Traders could submit jobs directly from Excel
- **Quick Turnaround**: Results typically returned within 15 minutes for smaller runs
- **Self-Service**: Traders could access the compute infrastructure independently

## Operational Challenges

### Azure Batch Reliability Issues
- **Frequent Updates**: Microsoft team in Redmond regularly introduced changes, often on Friday nights
- **Regional Rollouts**: Updates were phased across regions, causing inconsistent behavior
- **Common Problems**: Dead VMs, unexpected system behavior at scale
- **Mitigation Strategy**: Dual-region setup allowed failover when one region experienced issues

### Support Requirements
- **24/7 Monitoring**: Team received alerts at 3 AM when systems failed
- **Quick Recovery**: Could switch to backup region and restart failed jobs
- **Time Pressure**: Everything needed to complete before 7 AM business start

This system represents a sophisticated financial risk management platform that leveraged cloud elasticity to handle complex mathematical computations while providing both automated overnight processing and on-demand analytical capabilities for traders.
