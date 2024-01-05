using System.Diagnostics;

namespace BlazInfra.Shared;

public static class Telemetry
{
    public static readonly ActivitySource AppActivitySource = new("BlazInfra");
}